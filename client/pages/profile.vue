<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <div class="flex-wrapper">
      <div class="profile-image">
        <img src="/profileImage.jpg" />
      </div>
      <!-- /.profile-image -->
      <div class="profile-explanation">
        <div class="profile-name">
          <span>コットン</span>
        </div>
        <!-- /.profile-name -->
        <p>東京の某大学の学部生です。webエンジニアになるため日々勉強をしています。個人開発だけではなくチーム開発でプロジェクトマネージャーも務めています。</p>
        <div class="profile-link">
          <div><a href="https://github.com/cotton-alta"><img src="/github.png" /></a></div>
          <div><a href="https://twitter.com/cotton_alta_"><img src="/twitter.png" /></a></div>
          <div><a href="mailto:chinoknct@gmail.com"><img src="/mail.png" /></a></div>
        </div>
        <!-- /.profile-link -->
      </div>
      <!-- /.profile-explanation -->
    </div>
    <!-- /.flex-wraapper -->
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"

export default {
  components: { tabTitle },
  data() {
    return { pageTitle: "PROFILE" }
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

.flex-wrapper {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  @include mq {
    width: 90%;
    margin: 0 auto;
    flex-direction: row;
  }
}

.profile {
  &-image {
    width: 40%;
    text-align: center;
    margin: 20px 0px;
    @include mq {
      width: 30%;
      text-align: left;
    }
    img {
      width: 80%;
      border-radius: 50%;
    }
  }
  &-explanation {
    color: $string-color;
    width: 90%;
    line-height: 30px;
    @include mq {
      display: inline-block;
      width: 70%;
    }
    img {
      width: 2.5rem;
      margin: 0px 25px;
    }
  }
  &-name {
    margin: 20px 0px 15px 0px;
    text-align: center;
    span {
      font-size: calc(20px + 0.5vw);
      font-weight: bolder;
    }
  }
  &-link {
    margin: 20px 30px;
    display: flex;
    flex-direction: row;
    justify-content: center;
  }
}

</style>