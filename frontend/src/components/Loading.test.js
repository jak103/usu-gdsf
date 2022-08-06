import { mount } from '@vue/test-utils'
import Loading from './Loading.vue'

let wrapper;

// Occurs before each test.
beforeEach(() => {
    wrapper = mount(Loading, {
        data() {
            return {
                text: "Test Game",
            }
        }
    })
})

test('renders loading information', () => {
    const container = wrapper.get('[data-test="container"]');
    const text = wrapper.get('[data-test="text-wrapper"]');
    const progress = wrapper.get('[data-test="progress-wrapper"]');

    expect(container).toBeTruthy();
    expect(text).toBeTruthy();
    expect(text.text()).toBe("Loading Test Game");
    expect(progress).toBeTruthy();
})