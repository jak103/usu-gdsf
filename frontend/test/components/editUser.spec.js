import { shallowMount } from '@vue/test-utils'
import EditUser from "../../src/views/editUser.vue"


describe('Form Submit Working', () => {
    test('trigger', async () => {
        const wrapper = shallowMount(EditUser)

        // trigger the element
        await wrapper.find('v-btn').trigger('click')
      
        // assert some action has been performed, like an emitted event.
        expect(wrapper.emitted()).toHaveProperty('submit')
      })
})
