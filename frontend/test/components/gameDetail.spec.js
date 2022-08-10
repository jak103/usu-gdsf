import { shallowMount } from '@vue/test-utils'
import GameDetail from '../../src/components/GameDetail.vue'

describe('Game Detail default data rendering', () => {
  const tempData = {
    title: 'Game Title',
    description: 'A long Game Description',
    shortDescription: 'short descript',
    rating: 3,
    creationDate: '2022-01-01',
    version: 1.0,
    image: ''
  };

  const wrapper = shallowMount(GameDetail, {
    propsData: {
      gameId: 1
    }
  });
  test('title displays', () => {
    expect(wrapper.props().gameId).toBe(1);
  });
});