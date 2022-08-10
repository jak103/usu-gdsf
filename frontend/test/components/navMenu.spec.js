import {shallowMount} from '@vue/test-utils'
import NavMenu from '../../src/components/NavMenu.vue'

// tests by Kailey Bales


// test the navigation menu component
describe('NavMenu.vue', () => {

  const wrapper = shallowMount(NavMenu);

  test('renders a navigation menu', () => {
    //check if the navigation menu is rendered
    expect(wrapper.find('v-list-item').text()).toBe('Home')
  });
  test('renders a navigation menu with links', () => {
    //check if the navigation menu is rendered
    expect(wrapper.find('v-list-item').exists()).toBe(true)
  });
  test('home link goes to home page', () => {
    //check if clicking the home link goes to the home page
    expect(wrapper.find('v-list-item').text()).toContain('Home')
    expect(wrapper.find('v-list-item').attributes('to')).toBe('/')
  });


});
