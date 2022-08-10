import { mount  } from '@vue/test-utils'
import GameManagementPage from './GameManagementPage.vue'

test('Verify close new game leaves empty data in selectGame', () => {
    let wrapper = mount(GameManagementPage, {
        data() {
            return {
                games: [],
                showCreate: true,
                showEdit: false,
                selectedGame: {
                    name: "t",
                    developer: "e",
                    version: "s",
                    description: "t",
                    imagePath: "i",
                    downloadLink: "n"
                },
                error: ''
            }
        }, 

    })

    const cancelB = wrapper.find('#cancel')
    cancelB.trigger('click')

    expect(wrapper.vm.selectedGame.name).toBe(' ')
    expect(wrapper.vm.selectedGame.developer).toBe(' ')
    expect(wrapper.vm.selectedGame.version).toBe(' ')
    expect(wrapper.vm.selectedGame.description).toBe(' ')
    expect(wrapper.vm.selectedGame.imagePath).toBe(' ')
    expect(wrapper.vm.selectedGame.downloadLink).toBe(' ')
})

test('Verify close leaves empty data in selectGame', () => {
    let wrapper = mount(GameManagementPage, {
        data() {
            return {
                games: [],
                showCreate: false,
                showEdit: true,
                selectedGame: {
                    name: "t",
                    developer: "e",
                    version: "s",
                    description: "t",
                    imagePath: "i",
                    downloadLink: "n"
                },
                error: ''
            }
        }, 

    })

    const cancelB = wrapper.find('#cancel')
    cancelB.trigger('click')

    expect(wrapper.vm.selectedGame.name).toBe(' ')
    expect(wrapper.vm.selectedGame.developer).toBe(' ')
    expect(wrapper.vm.selectedGame.version).toBe(' ')
    expect(wrapper.vm.selectedGame.description).toBe(' ')
    expect(wrapper.vm.selectedGame.imagePath).toBe(' ')
    expect(wrapper.vm.selectedGame.downloadLink).toBe(' ')
})