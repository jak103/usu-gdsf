import { mount } from '@vue/test-utils'
import AdminAllUsers from '@/views/AdminAllUsers.vue'


describe('All Users prop rendering', () => {
    const testHeaders = [
        { text: 'ID', value: 'ID' },
        { text: 'Username', value: 'Username' },
        { text: 'Name', value: 'Displayname' },
        { text: 'Role', value: 'Role' }
    ];
    const testUsers = [
        {
            ID: 1,
            Username: 'test',
            Displayname: 'Test User',
            Role: 'Admin'
        }
    ];
    const wrapper = mount(AdminAllUsers)
    test('All Users rendering', () => {
        const id = wrapper.get('[data-test="id"]');
        const userName = wrapper.get('[data-test="username"]');
        const name = wrapper.get('[data-test="name"]');
        const role = wrapper.get('[data-test="role"]');
        expect (id).toEqual(testUsers[0].ID);
        expect (userName).toEqual(testUsers[0].Username);
        expect (name).toEqual(testUsers[0].Displayname);
        expect (role).toEqual(testUsers[0].Role);
    });
    
})