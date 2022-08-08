import { mount } from '@vue/test-utils'
import LoginPage from './LoginPage.vue'

test('redirects when logged in', () => {
  const wrapper = mount(LoginPage, {
    input: {
      username: 'admin',
      password: 'adminpass'
    }
  })
  
  const button = wrapper.find('button')
  button.trigger('click')
  expect(window.location.href).toBe('http://localhost/')
})