import {mount} from '@vue/test-utils'
import NavMenu from '@/components/navMenu.vue'



// test the navigation menu component
describe('NavMenu.vue', () => {
  let wrapper;
  beforeEach(() => {
    wrapper = mount(NavMenu);
  });
  test('renders a navigation menu', () => {
    //check if the navigation menu is rendered
    expect(wrapper.find('v-list').exists()).toBe(true)
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
