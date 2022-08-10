import { mount, flushPromises  } from '@vue/test-utils'
import * as gamesServices from '../services/gamesServices'
import AllGamesPage from './AllGamesPage.vue'

// Object to contain component.
let wrapper;

// Mock router
const mockRouter = {
    push: jest.fn()
}

// Occurs before each test.
beforeEach(() => {
    wrapper = mount(AllGamesPage, {
        data() {
            return {
                allGames: [],
                dataLoading: false
            }
        },
        created() {
            // Does nothing, prevents auto
            // kickoff of data retrieval
        },
        global:{
            mocks: {
                $router: mockRouter
            }
        }
    })
})

test('renders data table after loading', async () => {
    await wrapper.setData({
        dataLoading: false,
        allGames: [{}, {}, {}]
    });

    const title = wrapper.get('[data-test="title"]')
    expect(title).toBeTruthy()
    expect(title.text()).toBe("All Games")

    const table = wrapper.get('[data-test="data-table"]')
    expect(table).toBeTruthy()
    
    const pagination = wrapper.get('[data-test="pagination"]')
    expect(pagination).toBeTruthy()

    expect(wrapper.find('[data-test="loadbar"]').exists()).not.toBeTruthy()
});

test("renders load bar on load", async () => {
    await wrapper.setData({
        dataLoading: true,
        allGames: null
    });

    expect(wrapper.find('[data-test="title"]').exists()).not.toBeTruthy()
    expect(wrapper.find('[data-test="data-table"]').exists()).not.toBeTruthy()
    expect(wrapper.find('[data-test="pagination"]').exists()).not.toBeTruthy()

    expect(wrapper.find('[data-test="loadbar"]').exists()).toBeTruthy()
})

test("renders no data on empty data", async () => {
    await wrapper.setData({
        dataLoading: false,
        allGames: []
    });

    expect(wrapper.find('[data-test="title"]').exists()).not.toBeTruthy()
    expect(wrapper.find('[data-test="data-table"]').exists()).not.toBeTruthy()
    expect(wrapper.find('[data-test="pagination"]').exists()).not.toBeTruthy()
    expect(wrapper.find('[data-test="loadbar"]').exists()).not.toBeTruthy()

    expect(wrapper.find('[data-test="noData"]').exists()).toBeTruthy()
})

test("converts date to string", () => {
    const testDate = new Date()
    const results = wrapper.vm.getDateString(testDate)
    expect(results).toBe(
        `${testDate.getMonth()}/${testDate.getDate()}/${testDate.getFullYear()}`)
})

test("routes to game page", () => {
    const testId = "TESTID"
    const testName = "TestName"

    wrapper.vm.handleClickGame(testId, testName)

    expect(mockRouter.push).toHaveBeenCalledTimes(1)
    expect(mockRouter.push).toHaveBeenCalledWith(`/games/info/${testName}/${testId}`)
})

test("hits backend to get games list", async () => {
    const gameName = "Test Game"
    jest.spyOn(gamesServices, 'getAllGames').mockResolvedValue({ data: [{ Name: gameName }] })

    wrapper.vm.getGames()

    expect(gamesServices.getAllGames).toHaveBeenCalledTimes(1)

    await flushPromises()

    expect(wrapper.vm.allGames.length).toBe(1)
    expect(wrapper.vm.allGames[0].Name).toBe(gameName)
})