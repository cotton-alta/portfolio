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
        <label for="link" class="form-label">作品リンク</label>
        <div class="form-text">
          <input type="text" name="link" v-model="link" />
        </div>
        <!-- /.form-text -->
        <label for="github" class="form-label">git hub</label>
        <div class="form-text">
          <input type="text" name="github" v-model="github" />
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
  middleware({ store, redirect }) {
    if(!store.getters["isAuthenticated"]) {
      redirect('/')
    }
  },
  components: {
    tabTitle
  },
  data() {
    return { 
      pageTitle: "作品投稿",
      title: "",
      detail: "",
      image: "",
      link: "",
      github: ""
    }
  },
  methods: {
    uploadImage: function() {
      this.image = event.target.files[0]
    },
    postForm: function() {
      let formData = new FormData(),
          escapeDetail = this.detail.replace(/\n/g, '<br>')
      formData.append('title', this.title)
      formData.append('detail', escapeDetail)
      formData.append('link', this.link)
      formData.append('github', this.github)
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
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
  min-height: calc(100vh - 50px);
  @include mq {
    min-height: calc(100vh - 80px);
  }
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
    margin-bottom: 30px;
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