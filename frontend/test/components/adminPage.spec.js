import { mount } from '@vue/test-utils'
import AdminPage from "../../src/views/AdminPage.vue"

describe('HomeView Rendering', () => {

    const wrapper = mount(AdminPage);


    test('renders a form with text', () => {
        //check if the form is rendered
        expect(wrapper.find('This is where you login. Thanks for stopping by').exists()).toBe(true)
    }),

    test('renders a form with rows', () => {
        //check if the form is rendered
        expect(wrapper.find('AdminPage').exists()).toBe(true)
      })

})