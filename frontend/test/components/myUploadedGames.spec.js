import { shallowMount } from '@vue/test-utils'
import MyUploadedGames from "../../src/views/MyUploadedGames.vue"

describe('MyUploadedGames Testing', () =>{
    const testTitle = 'Game Info';
    const testAuthor = 'Game '
    const testDescription = 'Game'
    const wrapper = mount(MyUploadedGames);


    test('Renders a find on the text', () => {
        //check if the form is rendered
        expect(wrapper.find('All of my games').exists()).toBe(true)
      })
})