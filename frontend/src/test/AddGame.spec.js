import {mount} from '@vue/test-utils'
import AddGame from '@/views/AddGame.vue'

// tests by Kailey Bales

describe('AddGame.vue', () => {
  let wrapper;
  beforeEach(() => {
    wrapper = mount(AddGame);
  });

  test('renders a form', () => {
    //check if the form is rendered
    expect(wrapper.find('v-card').exists()).toBe(true)
  }),
  test('renders a form with inputs', () => {
    //check if the form is rendered
    expect(wrapper.find('input').exists()).toBe(true)
  }),
  test('renders a form with submit button', () => {
    //check if the form is rendered
    expect(wrapper.find('v-btn').exists()).toBe(true)
  }),
  test('sets the values', async () => {
    const input = wrapper.find('input')
  
    await input.setValue('test name')

  
    expect(input.element.value).toBe('test name')
  })
})
