describe('AboutView.vue', () => {
    const localVue = createLocalVue()
    let vuetify
  
    beforeEach(() => {
      vuetify = new Vuetify()
    })
  
    const mountFunction = options => {
      return mount(AboutView, {
        localVue,
        vuetify,
        ...options,
      })
    }
  
    it('should have a custom title and match snapshot', () => {
      const wrapper = mountFunction({
        propsData: { title: 'Fizzbuzz' },
      })
  
      expect(wrapper.html()).toMatchSnapshot()
    })
  })