import { mount } from '@vue/test-utils'
import form from './UserForm.vue'


test('Checks admin creation form action buttons', () => {
  const wrapper = mount(form, {
    propsData: {
      selectedUser : {
        firstName: 'Test',
        lastName: 'User',
        email: 'testUser@example.com',
        dob: '09-01-2000',
      },
      showSelf: true,
      isAdminCreation: true,
      newPassword: '',
    }
  })

  const adminCreateBtn = wrapper.find('#adminCreateBtn')
  const closeBtn = wrapper.find('#closeBtn')
  expect(adminCreateBtn.isVisible())
  expect(closeBtn.isVisible())
})

test('Checks admin creation form action buttons pt. 2', () => {
  const wrapper = mount(form, {
    propsData: {
      selectedUser : {
        firstName: 'Test',
        lastName: 'User',
        email: 'testUser@example.com',
        dob: '09-01-2000',
      },
      showSelf: true,
      isAdminCreation: false,  // This makes it a user edit form with different action buttons
      newPassword: '',
    }
  })
  const saveEditBtn = wrapper.find('#adminCreateBtn')
  const deleteUserBtn = wrapper.find('#deleteUserBtn')

  expect(() => {
    saveEditBtn.isVisible()
    deleteUserBtn.isVisible()
  }).toThrow('Cannot call isVisible on an empty DOMWrapper.')
})

test('Checks edit user form action buttons', () => {
  const wrapper = mount(form, {
    propsData: {
      selectedUser : {
        firstName: 'Test',
        lastName: 'User',
        email: 'testUser@example.com',
        dob: '09-01-2000',
      },
      showSelf: true,
      isAdminCreation: false,  // This makes it a user edit form with different action buttons
      newPassword: '',
    }
  })

  const saveEditBtn = wrapper.find('#saveEditBtn')
  const deleteUserBtn = wrapper.find('#deleteUserBtn')
  const closeBtn = wrapper.find('#closeBtn')
  expect(saveEditBtn.isVisible())
  expect(deleteUserBtn.isVisible())
  expect(closeBtn.isVisible())
})

test('Checks edit user form action buttons pt. 2', () => {
  const wrapper = mount(form, {
    propsData: {
      selectedUser : {
        firstName: 'Test',
        lastName: 'User',
        email: 'testUser@example.com',
        dob: '09-01-2000',
      },
      showSelf: true,
      isAdminCreation: false,  // This makes it a user edit form with different action buttons
      newPassword: '',
    }
  })

  const adminCreateBtn = wrapper.find('#adminCreateBtn')

  expect(() => {
    adminCreateBtn.isVisible()
  }).toThrow('Cannot call isVisible on an empty DOMWrapper.')
})

test('Test close form behavior', () => {
  const wrapper = mount(form, {
    propsData: {
      selectedUser : {
        firstName: 'Test',
        lastName: 'User',
        email: 'testUser@example.com',
        dob: '09-01-2000',
      },
      showSelf: true,
      isAdminCreation: false,
      newPassword: '',
    }
  })

  const closeBtn = wrapper.find('#closeBtn')
  closeBtn.trigger('click')

  const formElement = wrapper.find('v-form')
  expect(!formElement.isVisible())
  
})