import { mount } from '@vue/test-utils'
import carousel from './GameCarousel.vue'

test('renders Empty Game Carousel', () => {
  const wrapper = mount(carousel, {
    propsData: {
        screenshotUrls: []
    }
  })
  const carouselItems = wrapper.findAll('v-carousel-item')
  expect(carouselItems.length).toBe(0)
})

test('renders single URL Game Carousel', () => {
    const wrapper = mount(carousel, {
      data() { 
        return {
          screenshotUrls: ['https://i.ytimg.com/vi/-Gq2S0AXjNw/mqdefault.jpg']
        }
      }
    })
    const carouselItems = wrapper.findAll('v-carousel-item')
    expect(carouselItems.length).toBe(1)
  })