<template>
  <div>
    <van-divider>
      评论
    </van-divider>
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
        <div
          v-for="item in comments"
          :key="item.id + 'comment'"
          v-long-press="300"
          @long-press-start="onLongPressStart(item.id, item.content)"
          @long-press-stop="onLongPressStop(item.id, item.content)"
        >
          <van-cell :title="item.content">
            <template #label>
              <span>{{ item.create_time.substr(0, 19) }}</span>
            </template>
            <template #right-icon>
              <van-icon
                v-if="$store.state.user.uid == item.user_id"
                name="delete"
                @click="deleteComment(item.id, item.content)"
                style="line-height: inherit;"
              />
            </template>
          </van-cell>
        </div>
      </van-list>
      <van-divider />
    </div>

    <van-dialog
      v-model="show"
      title="修改评论内容"
      @confirm="editCommentConfirm"
      show-cancel-button
    >
      <van-field v-model="editCommentContent" label="评论" />
    </van-dialog>
  </div>
</template>

<script>
import Vue from "vue";
import LongPress from "vue-directive-long-press";

import { Field, Button, Divider, List, Cell, Icon } from "vant";
import { request } from "../../nework/request.js";
import { Dialog } from "vant";

// 全局注册
Vue.use(Dialog);
Vue.use(Field);
Vue.use(Icon);
Vue.use(Divider);
Vue.use(Button);
Vue.use(List);
Vue.use(Cell);
export default {
  name: "TravelNoteComment",
  directives: {
    "long-press": LongPress,
  },
  data() {
    return {
      content: "",
      comments: [],
      loading: false,
      finished: false,
      show: false,
      editCommentContent: "",
      editCommentId: 0,
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
          if (res.data > 0) {
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
    onLongPressStart(id, content) {
      // triggers after 300ms of mousedown
      console.log("long press: ", id);
      console.log("long press content: ", content);
      this.editCommentContent = content;
      this.editCommentId = id;
    },
    onLongPressStop(id, content) {
      // triggers on mouseup of document
      console.log("long end press: ", id);
      console.log("long end press content: ", content);
      this.show = true;
      this.editCommentContent = content;
      this.editCommentId = id;
    },
    editCommentConfirm() {
      request({
        url: "comment/update",
        method: "post",
        params: {
          comment_id: this.editCommentId,
          content: this.editCommentContent,
        },
      }).then((res) => {
        console.log("edit comment res: ", res);
        this.getCommentsData();
        this.editCommentContent = "";
        this.editCommentId = 0;
      });
      console.log("confirm:", this.editCommentContent);
    },
  },
};
</script>

<style></style>
