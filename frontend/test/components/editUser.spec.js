import { shallowMount } from '@vue/test-utils'
import EditUser from "../../src/views/EditUser.vue"


describe('Edit user', () => {
    test('text field exists', async () => {
        const wrapper = shallowMount(EditUser)

        expect(wrapper.find('v-text-field').exists()).toBe(true)
      })
})
