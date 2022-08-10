import {  mount } from '@vue/test-utils'
import menu from './NavMenu.vue'

test('renders Nav Menu', () => {
  const wrapper = mount(menu)
  const title = wrapper.get('[data-test="listItem1"]')
  expect(title.text()).toBe("Game Dev Store")
})

test('Checks Nav Menu Items', () => {
  const wrapper = mount(menu)
})
