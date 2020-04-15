<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <div class="form-wrapper">
      <label for="email" class="form-label">連絡先メールアドレス</label>
      <div class="form-text">
        <input type="text" name="email" v-model="email" />
      </div>
      <!-- /.form-text -->
      <label for="content" class="form-label">お問い合わせ内容</label>
      <div class="form-text">
        <textarea name="content" rows="3" v-model="content"></textarea>
      </div>
      <!-- /.form-text -->
      <div class="form-submit"><input type="submit" value="送信"/></div>
      <div class="form-submit" @click="sendMessage">送信</div>
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
      pageTitle: "CONTACT" ,
      email: "",
      content: ""
    }
  },
  methods: {
    sendMessage: function() {
      axios.post("/api/contact"

      )
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