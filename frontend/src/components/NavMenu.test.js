import { shallowMount, mount } from '@vue/test-utils'
import menu from './NavMenu.vue'

test('renders Nav Menu', () => {
  const wrapper = mount(menu)
  const todo = wrapper.get('[data-test="listItem1"]')
  expect(todo.text()).toBe("Game Dev Store")
})