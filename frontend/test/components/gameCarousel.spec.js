import { mount } from '@vue/test-utils'
import GameCarousel from "../../src/components/GameCarousel.vue"

describe('GameCarousel renders cards', () => {
    const testColors = ['primary','secondary','info','warning','white','teal', 'watermelon']

    const wrapper = mount(GameCarousel)

    test('Carousel Rendering', () => {
        const color = wrapper.get('[data-test="test"]')
        expect(color.text()).toBe("primary")
    })
});