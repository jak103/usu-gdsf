import {mount} from '@vue/test-utils'
import AboutView from '@/AboutView'

describe('AboutView', () => {
    test('initial state false', () => {
        const wrapper = mount(AboutView)

        expect(wrapper.html()).toContain(false)
    })
    test('when clicked, show changes to true', () => {
        const wrapper = mount(AboutView)
        const button = wrapper.find('button')

        expect(wrapper.html().toContain(true))
    })
})