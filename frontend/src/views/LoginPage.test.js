import { mount } from '@vue/test-utils'
import LoginPage from './LoginPage.vue'

let wrapper;
const testInput = {
  username: "admin",
  password: "adminpass"
}

test('working log in', () => {
  wrapper = mount(LoginPage)
  const adminName = "admin"
  const adminPassword = "adminpass"
})