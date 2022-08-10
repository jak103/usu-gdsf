import { shallowMount } from '@vue/test-utils'
import HomeView from "../../src/views/HomeView.vue"

describe('HomeView Rendering', () =>{
    const testTitle = 'HomeView';
    const testAuthor = 'Game Author'
    const testDescription = 'Game Description'
    const wrapper = shallowMount( HomeView, {
        props : {
            name : 'HomeView',
            // author : testAuthor,
            // description : testDescription,
        },
    });

    test('Title', () => {
        // expect(wrapper.props().name).toBe(testTitle);
    });
})