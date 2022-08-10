import { shallowMount } from '@vue/test-utils'
import GameList from "../../src/components/GameList.vue"

describe('GameList Testing', () =>{
    const testTitle = 'Test Card Title';
    const testAuthor = 'Game Author'
    const testDescription = 'Game Description'
    const wrapper = shallowMount( GameList, {
        propsData : {
            title : testTitle,
            author : testAuthor,
            description : testDescription,
        },
    });

    test('Title', () => {
        expect(wrapper.props().title).toBe(testTitle);
    });
    test('Author', () => {
        expect(wrapper.props().author).toBe(testAuthor);
    });
    test('Author', () => {
        expect(wrapper.props().author).toBe(testAuthor);
    });
})