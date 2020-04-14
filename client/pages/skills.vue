<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <div class="logo-wrapper">
      <skill class="logo-card" :skill="skill" v-for="skill in skills" :key="skill.id"/>
    </div>
    <!-- /.logo-wrapper -->
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"
import skill from "~/components/ui/skill.vue"
import apache from "~/static/logo/apache.png"
import centos from "~/static/logo/centos.png"
import css from "~/static/logo/css.png"
import django from "~/static/logo/django.png"
import html from "~/static/logo/html.png"
import javascript from "~/static/logo/javascript.png"
import mongo from "~/static/logo/mongo.png"
import node from "~/static/logo/node.png"
import nuxt from "~/static/logo/nuxt.png"
import python from "~/static/logo/python.png"
import vue from "~/static/logo/vue.png"

export default {
  components: {
    tabTitle,
    skill
  },
  data() {
    return {
      pageTitle: "SKILLS",
      skills: {
        apache: { name: "apache", img: apache, years: "1年" },
        centos: { name: "centos", img: centos, years: "1年" },
        css: { name: "css", img: css, years: "1年" },
        django: { name: "django", img: django, years: "1年" },
        html: { name: "html", img: html, years: "1年" },
        javascript: { name: "javascript", img: javascript, years: "1年" },
        mongo: { name: "mongo", img: mongo, years: "1年" },
        node: { name: "node", img: node, years: "1年" },
        nuxt: { name: "nuxt", img: nuxt, years: "1年" },
      }
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

.logo {
  &-wrapper {
    margin: 0 auto;
    display: flex;
    flex-direction: row;
    justify-content: center;
    flex-wrap: wrap;
    @include mq {
      width: 600px;
      justify-content: space-between;
    }
  }
  &-card {
    margin: 20px 15px;
    @include mq(sm) {
      margin: 20px 30px;
    }
  }
}

</style>