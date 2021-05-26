FROM nginx
COPY test.go /user/src/app
ADD . /user/src/app
# 如果源文件是一个压缩文件，ADD会自动解压缩这个压缩文件到目标路径去
# Docker 官方的最佳实践文档中要求尽可能的使用COPY、因为COPY的语义很明确，就是复制文件而已 而ADD则包含了更复杂的功能。
# 因此在 COPY 和 ADD 指令中选择的时候，可以遵循这样的原则，所有的文件复制均使用 COPY 指令，仅在需要自动解压缩的场合使用 ADD。
ADD test.go /
ENV HOME /user
# Docker不是虚拟机，容器就是进程，那么在启动容器的时候，需要指定所运行的程序和参数。
CMD echo $HOME
# 入口点的目的和CMD是一样的 https://www.golangroadmap.com/class/docker/image/dockerfile/entrypoint.html
ENTRYPOINT []
# 设置环境变量
ENV TEST test
# 构建参数 和ENV的效果一样，都是设置环境变量，所不同的是ARG设置的环境变量在将来容器运行时是不会存在这些环境变量的。因此不要用ARG设置密码
# 之类的信息
ARG NOPASSWORD=123


VOLUME /data
USER root

HEALTHCHECK --interval=5s --timeout=3s \
CMD curl -fs http://localhost/ || exit 1


# 添加元数据
LABEL author="joy"


