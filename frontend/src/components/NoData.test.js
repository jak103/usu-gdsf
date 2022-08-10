import { mount } from '@vue/test-utils'
import NoData from './NoData.vue'

let wrapper;

// Mock router
const mockRouter = {
    push: jest.fn()
}

// Occurs before each test.
beforeEach(() => {
    wrapper = mount(NoData, {
        data() {
            return {
                text: "Test Game"
            }
        },
        global:{
            mocks: {
                $router: mockRouter
            }
        }
    })
})

test('renders no data information', () => {
    const container = wrapper.get('[data-test="container"]');
    const text = wrapper.get('[data-test="text-wrapper"]');
    const homeBtn = wrapper.get('[data-test="home-button"]');

    expect(container).toBeTruthy();
    expect(text).toBeTruthy();
    expect(text.text()).toBe("Oops! There was no data for: Test Game");
    expect(homeBtn).toBeTruthy();
})

test("routes to home page on button click", () => {
    const homeBtn = wrapper.get('[data-test="home-button"]');
    
    homeBtn.trigger('click.stop');

    expect(mockRouter.push).toHaveBeenCalledTimes(1)
    expect(mockRouter.push).toHaveBeenCalledWith("/")
})