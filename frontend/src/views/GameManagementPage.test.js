import { mount, flushPromises  } from '@vue/test-utils'
import axios from 'axios'
import GameManagementPage from './GameManagementPage.vue'

test('Verify close new game leaves empty data in selectGame', () => {
    const wrapper = mount(GameManagementPage, {
        propsData: {
            games: [],
			showCreate: true,
			showEdit: ref(false),
			selectedGame: ref({
				name: "t",
				developer: "e",
				version: "s",
				description: "t",
				imagePath: "i",
				downloadLink: "n"
			}),
			error: ''
        }, 

    })

    const cancelB = wrapper.find('#cancel')
    cancelB.trigger('click')

    expect(wrapper.props().selectedGame.name).toBe(' ')
    expect(wrapper.props().selectedGame.developer).toBe(' ')
    expect(wrapper.props().selectedGame.version).toBe(' ')
    expect(wrapper.props().selectedGame.description).toBe(' ')
    expect(wrapper.props().selectedGame.imagePath).toBe(' ')
    expect(wrapper.props().selectedGame.downloadLink).toBe(' ')
})

test('Verify close leaves empty data in selectGame', () => {
    const wrapper = mount(GameManagementPage, {
        propsData: {
            games: [],
			showCreate: ref(false),
			showEdit: true,
			selectedGame: ref({
				name: "t",
				developer: "e",
				version: "s",
				description: "t",
				imagePath: "i",
				downloadLink: "n"
			}),
			error: ''
        }, 

    })

    const cancelB = wrapper.find('#cancel')
    cancelB.trigger('click')

    expect(wrapper.props().selectedGame.name).toBe(' ')
    expect(wrapper.props().selectedGame.developer).toBe(' ')
    expect(wrapper.props().selectedGame.version).toBe(' ')
    expect(wrapper.props().selectedGame.description).toBe(' ')
    expect(wrapper.props().selectedGame.imagePath).toBe(' ')
    expect(wrapper.props().selectedGame.downloadLink).toBe(' ')
})