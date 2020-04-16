<template>
  <div class="button-wrapper">
    <div class="flex-image"><img src="/logo/twitterbtn.png" /></div>
    <div class="flex-text"><span @click="windowOpen(twitterLink)">Tweet</span></div>
  </div>
  <!-- /.button-wrapper -->
</template>

<script>
export default {
  props: ["title"],
  data() {
    let path = encodeURIComponent("https://cotton-portfolio.org") + `/articles/${this.title}`,
        text = "[" + this.title + "]"
    return {
      twitterLink: `https://twitter.com/intent/tweet?url=${path}&text=${text}&lang=ja`
    }
  },
  methods: {
    windowOpen: function(link) {
      return window.open(link, '_blank', 'top=100,left=100,width=600,height=500')
    }
  }
}
</script>

<style lang="scss" scoped>

.button-wrapper {
  margin: 0 auto;
  width: 85px;
  height: 30px;
  border-radius: 4px;
  color: #ffffff;
  background-color: #1B95E0;
  cursor: pointer;
  display: flex;
  flex-direction: row;
  justify-content: center;
}

.flex {
  &-image {
    height: 18px;
    img {
      height: 18px;
      margin: 6px 3px 0px 0px;
    }
  }
  &-text {
    padding-top: 1px;
    margin-right: 2px;
  }
}

</style>