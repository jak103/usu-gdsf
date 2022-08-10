import { shallowMount } from '@vue/test-utils'
import GameCarousel from "../../src/components/GameCarousel.vue"

describe('GameCarousel renders cards', () => {
    const wrapper = shallowMount(GameCarousel)

    test('Carousel Rendering', () => {
        expect(wrapper.findAll('[test-data="gameData"]')).toHaveLength(0)
    })
});