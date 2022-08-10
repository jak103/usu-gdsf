import { mount } from '@vue/test-utils'
import rating from './Rating.vue'

test('Rating component is visible', () => {
    const wrapper = mount(rating)
    const ratingElement = wrapper.get('[data-test="rating-main"]')
    expect(ratingElement.isVisible());
})

test('renders rating componenet', () => {
    const wrapper = mount(rating)

      const vRating = wrapper.findAll('v-rating')
      expect(vRating.length).toBe(1)
    })

test('rendering starValue of Rating.vue', async () => {
    const wrapper = mount(rating, {
        propsData: {
            starValue:3,
        }
      })
    const ratingValue = wrapper.find('[data-rating-test="rating"]')

    expect(ratingValue.html()).toContain('<span class="pr-1" data-rating-test="rating"> (3) </span>')
})

test('rendering on click for rating component', async () => {
    const wrapper = mount(rating)

    await wrapper.find('v-rating').trigger('click')

    expect(wrapper.emitted().click[0].length).toBe(1)
})


