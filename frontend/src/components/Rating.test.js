import { mount } from '@vue/test-utils'
import rating from './Rating.vue'

test('renders rating componenet', () => {
    const wrapper = mount(rating, {
        propsData: {
            starBackgroundColor:"#8A8D8F",
            isClearable:"false",
            closeDelay:"0",
            starColor:"#8A8D8F",
            isDark:"false",
            isDense:"true",
            isLight:"false",
            openDelay:"0",
            isReadOnly:"false",
            isRipple:"false",
            isHalfIncrements:"false",
            isHover:"false",
            isSamll:"false",
            starSize:"64",
            starValue:"3",
            isXLarge:"false",
            isXSmall:"false"
        }
      })
      const vRating = wrapper.findAll('v-rating')
      expect(vRating.length).toBe(1)
    })

test('rendering on click for rating', async () => {
    const wrapper = mount(rating)
    await wrapper.find('v-rating').trigger('click')
    expect(wapper.emitted().click[0][0]).toBe("MouseEvent")
    
})


