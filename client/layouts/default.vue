<template>
  <div>
    <headerBox id="top"/>
    <nuxt />
    <footerBox />
    <scrollButton v-show="show" class="scroll-button"/>
  </div>
</template>

<script>
import headerBox from "~/components/layouts/header.vue"
import footerBox from "~/components/layouts/footer.vue"
import scrollButton from "~/components/ui/topScrollButton.vue"

export default {
  data() {
    return { 
      show: false
    }
  },
  components: {
    headerBox,
    footerBox,
    scrollButton
  },
  mounted() {
    window.addEventListener("scroll", () => {
      this.show = (window.scrollY > 150)
    })
  }
}
</script>

<style>

html {
  font-family: 'Roboto', 'Noto Sans JP';
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

*,
*:before,
*:after {
  box-sizing: border-box;
  margin: 0;
}

.scroll-button {
  position: fixed;
  right: 5%;
  bottom: 5%;
}

</style>
