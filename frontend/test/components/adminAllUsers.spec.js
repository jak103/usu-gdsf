import { shallowMount } from '@vue/test-utils'
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
            id: 1,
            username: 'test',
            displayname: 'Test User',
            role: 'Admin'
        }
    ];
    const wrapper = shallowMount(AdminAllUsers, {
        data() {
            return {
                users: testUsers,
                headers: testHeaders
            }
        }
    })
    test('All Users Headers', () => {
        const id = wrapper.findAll('th')[0].text();
        const userName = wrapper.findAll('th')[1].text();
        const name = wrapper.findAll('th')[2].text();
        const role = wrapper.findAll('th')[3].text();
        expect(id).toBe('ID');
        expect(userName).toBe('Username');
        expect (name).toBe('Name');
        expect (role).toBe('Role');
    });

    test('All Users Rendering', () => {
        const id = wrapper.findAll('td')[0].text();
        const userName = wrapper.findAll('td')[1].text();
        const name = wrapper.findAll('td')[2].text();
        const role = wrapper.findAll('td')[3].text();
        expect(id).toBe('1');
        expect(userName).toBe('test');
        expect (name).toBe('Test User');
        expect (role).toBe('Admin');
    });

})