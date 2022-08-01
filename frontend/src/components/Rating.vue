<template>
		<v-card-actions> 
			<span class="pr-2">
				{{ rating }}
			</span>
			<v-rating
				v-model="rating"
                :background-color="starBackgroundColor" 
                :clearable="clearable"
                :close-delay="closeDelay"
				:color="starColor" 
                :dark="isDark"
                :dense="isDense"
                :light="isLight"
                :open-delay="openDelay"
                :readonly="isReadOnly"
                :ripple="isRipple"
                :size="starSize"
				:half-increments="isHalfIncrement"
                :hover="isHover"
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
        starSize:Number,
        isHalfIncrements:Boolean,
        isHover:Boolean,
        value:Number
    },

    data: () => ({
        rating: 3,
    }),
}
</script>