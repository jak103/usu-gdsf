import { mount } from '@vue/test-utils'
import rating from './Rating.vue'

test('Rating component is visible', () => {
    const wrapper = mount(rating)
    const ratingElement = wrapper.get('[data-test="rating-main"]')
    expect(ratingElement.isVisible());
})