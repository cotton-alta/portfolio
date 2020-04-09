<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
      <div class="form-wrapper">
        <label for="title" class="form-label">タイトル</label>
        <div class="form-text">
          <input type="text" name="title" v-model="title" />
        </div>
        <!-- /.form-text -->
        <label for="detail" class="form-label">詳細</label>
        <div class="form-text">
          <textarea name="detail" rows="3" v-model="detail"></textarea>
        </div>
        <!-- /.form-text -->
        <label for="image" class="form-label">画像</label>
        <div class="form-text">
          <input type="file" v-on:change="uploadImage" name="file" />
        </div>
        <!-- /.form-text -->
        <div @click="postForm" class="form-submit"><input type="submit" value="投稿"/></div>
      </div>
      <!-- /.form-wrapper -->
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"
import axios from "axios"

export default {
  components: {
    tabTitle
  },
  data() {
    return { 
      pageTitle: "作品投稿",
      title: "",
      detail: "",
      image: ""
    }
  },
  methods: {
    uploadImage: function() {
      this.image = event.target.files[0]
      console.log("OK!")
    },
    postForm: function() {
      let formData = new FormData()

      formData.append('title', this.title)
      formData.append('detail', this.detail)
      formData.append('file', this.image)
      axios.put("/api/works",
        formData,
        { header: { 'Content-Type': 'multipart/form-data' } }
      ).then(result => {
        console.log(result)
        this.$router.go({path: this.$router.currentRoute.path, force: true})
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  display: block;
  min-height: calc(100vh - 80px);
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
}

.form {
  &-wrapper {
    display: flex;
    width: 100%;
    flex-direction: column;
    align-items: center;
  }
  &-label {
    max-width: 480px;
    width: 75%;
    text-align: left;
    font-size: calc(15px + 0.3vw);
    font-weight: bolder;
    color: $menu-background;
    margin-top: 20px;
  }
  &-text {
    max-width: 500px;
    width: 80%;
    margin: 10px 0px;
    input {
      width: 100%;
      font-size: calc(15px + 0.3vw);
      line-height: 2rem;
      border-radius: 5px;
      display: inline-block;
    }
    textarea {
      width: 100%;
      font-size: calc(15px + 0.3vw);
      font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI',
        Roboto, 'Helvetica Neue', Arial, sans-serif;
      line-height: 2rem;
      border-radius: 5px;
      display: inline-block;
    }
  }
  &-submit {
    width: 60%;
    max-width: 300px;
    input {
      font-size: calc(15px + 0.3vw);
      font-weight: bolder;
      background-color: $menu-background;
      color: $main-background;
      width: 100%;
      margin-top: 40px;
      line-height: 2rem;
      border-radius: 5px;
      display: inline-block;
    }
  }
}
</style>