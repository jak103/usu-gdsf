import { mount } from '@vue/test-utils'
import HomeView from "../../src/views/HomeView.vue"

describe('HomeView Rendering', () => {
    const testTitle = 'HomeView';
    const testAuthor = 'Game Author'
    const testDescription = 'Game Description'

    const wrapper = mount(HomeView);


    test('renders a form with container', () => {
        //check if the form is rendered
        expect(wrapper.find('v-container').exists()).toBe(true)
    }),

    test('renders a form with rows', () => {
        //check if the form is rendered
        expect(wrapper.find('v-row').exists()).toBe(true)
      })

})