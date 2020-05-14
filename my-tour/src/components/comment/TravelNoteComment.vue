<template>
  <div>
    <van-divider />
    <!-- 输入评论 -->
    <div class="input_comment">
      <van-field
        v-model="content"
        center
        required
        autosize
        clearable
        type="textarea"
        placeholder="请输入评论"
      >
        <template #button>
          <van-button size="small" type="primary" @click="submitComment"
            >提交</van-button
          >
        </template>
      </van-field>
    </div>
    <div class="show_comments">
      <van-list
        v-model="loading"
        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <van-cell
          v-for="item in comments"
          :key="item.id + 'comment'"
          :title="item.content"
        >
          <template #label>
            <span>{{ item.create_time.substr(0, 19) }}</span>
          </template>
          <template #right-icon>
            <van-icon
              name="delete"
              @click="deleteComment(item.id)"
              style="line-height: inherit;"
            />
          </template>
        </van-cell>
      </van-list>
      <van-divider />
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import { Field, Button, Divider, List, Cell, Icon } from "vant";
import { request } from "../../nework/request.js";

Vue.use(Field);
Vue.use(Icon);
Vue.use(Divider);
Vue.use(Button);
Vue.use(List);
Vue.use(Cell);
export default {
  name: "TravelNoteComment",
  data() {
    return {
      content: "",
      comments: [],
      loading: false,
      finished: false,
    };
  },
  props: {
    travel_note_id: Number,
  },
  mounted() {
    this.getCommentsData();
  },
  methods: {
    onLoad() {
      this.getCommentsData();
      this.finished = true;
    },
    submitComment() {
      if (this.content != "") {
        request({
          url: "comment",
          method: "post",
          params: {
            travel_note_id: this.travel_note_id,
            content: this.content,
          },
        }).then((res) => {
          console.log("comment content submit req:", res);
          if (res.data == 1) {
            this.$toast.success("评论成功！");
            this.content = "";
          }
          this.getCommentsData();
        });
      } else {
        this.$toast.fail("输入不能为空！");
      }
    },
    deleteComment(id) {
      console.log("travel comment id: ", id);
      request({
        url: "comment/" + id,
        method: "DELETE",
      }).then((res) => {
        if (res.data > 0) {
          this.$toast.success("删除成功！");
          this.getCommentsData();
        }
      });
    },
    getCommentsData() {
      request({
        url: "comment",
        method: "get",
        params: {
          travel_note_id: this.travel_note_id,
        },
      }).then((res) => {
        console.log("comments req:", res);
        this.comments = res.data.comments;
      });
    },
  },
};
</script>

<style></style>
