import {mount} from '@vue/test-utils'
import AddGame from '@/views/AddGame.vue'

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
  }),
  test('test submitGame()', async() => {
    const input = wrapper.find('input')
    await input.setValue('test name')
    const button = wrapper.find('v-btn')
    await button.trigger('click')
    //check if gameName is set to test name
    expect(wrapper.vm.gameName).toBe('test name')
  }),
  test('test submitGame() again', async() => {
    const input = wrapper.find('#gameDesc')
    await input.setValue('test description')
    const button = wrapper.find('v-btn')
    await button.trigger('click')
    //check if gameName is set to test name
    expect(wrapper.vm.gameDesc).toBe('test description')
  })
})