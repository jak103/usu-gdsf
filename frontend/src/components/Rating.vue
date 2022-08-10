<template>
		<v-card-actions> 
			<span class="pr-1" data-rating-test="rating">
				({{ rating }})
			</span>
			<v-rating
        data-test="rating-main"
				v-model="rating"
        :background-color="starBackgroundColor"
        :clearable="isClearable"
        :close-delay="closeDelay"
				:color="starColor" 
        :dark="isDark"
        :density="isDense ? 'compact' : 'default'"
        :light="isLight"
        :open-delay="openDelay"
        :readonly="isReadOnly"
        :ripple="isRipple"
        :size="starSize"
				:half-increments="isHalfIncrements"
        :hover="isHover"
        :small="isSmall"
        :value="starValue"
        :x-large="isXLarge"
        :x-small="isXSmall"
			></v-rating> 
		</v-card-actions>
</template>

<script> 
export default {
    name: "Rating",

    props: {
        starBackgroundColor:String,
        isClearable:Boolean,
        closeDelay:Number,
        starColor:String,
        isDark:Boolean,
        isDense:Boolean,
        isLight:Boolean,
        openDelay:Number,
        isReadOnly:Boolean,
        isRipple:Boolean,
        isHalfIncrements:Boolean,
        isHover:Boolean,
        isSmall:Boolean,
        starSize:Number,
        starValue:Number,
        isXLarge:Boolean,
        isXSmall:Boolean,
        rating:Number
    },

    data: () => ({
      return: {
        rating: 4
      }
    }),
}
</script>