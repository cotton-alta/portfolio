<template>
  <div class="btn-wrapper" @click="deleteItem">
    <span>削除</span>
  </div>
  <!-- /.btn-wrapper -->
</template>

<script>
import axios from "axios"

export default {
  props: [ "title" ],
  methods: {
    deleteItem: function() {
      let item = encodeURIComponent(this.title)
      console.log(this.$route.path)
      if(this.$route.path === "/edit/worksList") {
        axios.delete(`/api/works/${item}`)
          .then(result => {
            console.log(result)
            this.$router.go({path: this.$router.currentRoute.path, force: true})
          })
      } else {
        axios.delete(`/api/articles/${item}`)
          .then(result => {
            console.log(result)
            this.$router.go({path: this.$router.currentRoute.path, force: true})
          })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.btn-wrapper {
  width: 100px;
  height: 40px;
  line-height: 40px;
  text-align: center;
  border-radius: 5px;
  background-color: #D0D0D2;
  span {
    display: inline-block;
    width: 100%;
    height: 100%;
    font-size: calc(15px + 0.1vw);
    font-weight: bolder;
    text-decoration: none;
    color: $main-background;
    cursor: pointer;
  }
}
</style>