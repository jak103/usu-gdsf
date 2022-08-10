/*
    Test that dummy data is properly detected from card creation.
    Once Backend is implemented, store values in a list to be exported. Modify test to see if propsData is populated.
*/
import { shallowMount } from '@vue/test-utils'
import GameList from "../../src/components/GameList.vue"

describe('GameList Card Rendering', () =>{
    const testTitle = 'Game Title';
    const testAuthor = 'Game Author'
    const testDescription = 'Game Description'
    const testImage = 'https://cdn.cloudflare.steamstatic.com/steam/apps/489830/capsule_616x353.jpg?t=1650909796';
    const wrapper = shallowMount( GameList, {
        propsData : {
            title : testTitle,
            author : testAuthor,
            description : testDescription,
            image : testImage
        },
    });

    test('Title', () => {
        expect(wrapper.props().title).toBe(testTitle);
    });
    test('Author', () => {
        expect(wrapper.props().author).toBe(testAuthor);
    });
    test('Description', () => {
        expect(wrapper.props().description).toBe(testDescription);
    });
    test('Image', () => {
        expect(wrapper.props().image).toBe(testImage);
    });
})