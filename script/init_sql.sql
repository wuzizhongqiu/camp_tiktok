SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
USE camps_tiktok;

DROP TABLE IF EXISTS `t_comment`;
CERATE TABLE `t_comment`
(
    `id`    bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论id，自增主键',
    `user_id` bigint(20) NOT NULL COMMENT '评论发布用户id',
    `video_id` bigint(20) NOT NULL COMMENT '评论视频id',
    `parent_id` bigint(20)  COMMENT '父评论的id',
    `parent_user_id` bigint(20)  COMMENT '父评论的user_id',
    `reply_id`  bigint(20)  COMMENT '回复评论的id,默认为二级评论',
    `reply_user_id` bigint(20)  COMMENT '回复评论的user_id,默认为二级评论',
    `comment_level` tinyint(4) NOT NULL COMMENT '评论的等级，1为一级评论，2为二级评论',
    `comment_text` varchar(255) NOT NULL COMMENT '评论内容',
    `create_time` datatime  NOT NULL COMMENT '评论发布时间',
     PRIMARY KEY (`id`),
     KEY `idx_video` (`video_id`,`comment_level`) USING BTREE,
     KEY `idx_create_time` (`created_time`) USING BTREE,
)