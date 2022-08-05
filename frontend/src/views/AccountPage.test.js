import { mount, flushPromises } from '@vue/test-utils'
import AccountPage from './AccountPage.vue'

let wrapper;
const mockUser = {
  Id: "123",
  email: "example@usu.edu",
  firstName: "John",
  lastName: "Doe",
  dateOfBirth: "Jan, 1, 2022",
  joinedTimestamp: 123456789,
  image: "https://randomuser.me/api/portraits/men/85.jpg"
}

test('renders Account Page', () => {
  wrapper = mount(AccountPage)
  const userImage = wrapper.get('[data-test="user-image"]')
  const userName = wrapper.get('[data-test="name"]')
  const email = wrapper.get('[data-test="email"]')
  const bday = wrapper.get('[data-test="bday"]')
  expect(userImage).toBeTruthy()
  expect(userName.text()).toEqual(`${mockUser.firstName} ${mockUser.lastName}`)
  expect(email.text()).toEqual(`${mockUser.email}`)
  expect(bday.text()).toEqual(`Birthday: ${mockUser.dateOfBirth}`)
})