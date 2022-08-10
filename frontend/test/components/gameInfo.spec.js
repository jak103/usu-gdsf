import { shallowMount } from '@vue/test-utils'
import GameInfo from "../../src/components/GameInfo.vue"

describe('GameInfo Testing', () =>{
    const testTitle = 'Game Info';
    const testAuthor = 'Game '
    const testDescription = 'Game'
    const wrapper = shallowMount( GameInfo, {
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