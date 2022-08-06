import { shallowMount } from '@vue/test-utils'
import GameCard from "../../src/components/GameCard.vue"


describe('GameCard prop rendering', () => {
    const testTitle = 'Test Card Title';
    const testAuthor = 'Jimmy John';
    const testDescription = 'Here is a description';
    const testImage = 'https://st2.depositphotos.com/2927537/7025/i/950/depositphotos_70253417-stock-photo-funny-monkey-with-a-red.jpg';
    const wrapper = shallowMount( GameCard, {
        propsData : {
            title : testTitle,
            author : testAuthor,
            description : testDescription,
            image : testImage
        },
    });
    test('title is present', () => {
        expect(wrapper.props().title).toBe(testTitle);
    });
    test('author is present', () => {
        expect(wrapper.props().author).toBe(testAuthor);
    });
    test('description is present', () => {
        expect(wrapper.props().description).toBe(testDescription);
    });
    test('image is present', () => {
        expect(wrapper.props().image).toBe(testImage);
    });
})
