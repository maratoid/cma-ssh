FROM centos:7.4.1708 as downloader

RUN yum update -y \
    && yum install epel-release -y \
    && yum install yum-utils wget -y \
    && yum clean all

# all baked kubernetes versions
ARG K8S_VERSIONS="1.10.6 1.11.2"
# helm versions
ARG HELM_VERSIONS="v2.11.0"
# docker version
ARG DOCKER_VERSION="1.13.1-75.git8633870.el7.centos"

# add docker repo, add kubernetes repo
RUN yum-config-manager --add-repo https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64 \
    && yum update -y

# pre create folders
RUN mkdir -p /var/tmp/yaml \
    && mkdir -p /var/tmp/bin \
    && mkdir -p /var/tmp/rpm

# donwload all k8s rpms and helm
RUN helmversions=($HELM_VERSIONS); k8sversions=($K8S_VERSIONS); for ver in ${versions[*]}; do mkdir -p /var/tmp/${ver}; done \
    && for ver in ${k8sversions[*]}; do yumdownloader --assumeyes --resolve --destdir /var/tmp/rpm/k8s-${ver} \
        kubectl-${ver} kubeadm-${ver} kubelet-${ver} kubernetes-cni; done \
    && yumdownloader --assumeyes --resolve --destdir /var/tmp/rpm/docker audit audit-libs libselinux libsepol libsemanage \
        device-mapper-persistent-data lvm2  conntrack docker-${DOCKER_VERSION} \
        docker-client-${DOCKER_VERSION} docker-common-${DOCKER_VERSION} \
    && yumdownloader --assumeyes --resolve --destdir /var/tmp/rpm/nginx nginx \
    && yumdownloader --assumeyes --resolve --destdir /var/tmp/rpm/curl curl \
    && yumdownloader --assumeyes --resolve --destdir /var/tmp/rpm/wget wget \
    && wget --directory-prefix=/var/tmp/yaml https://raw.githubusercontent.com/coreos/flannel/v0.10.0/Documentation/kube-flannel.yml \
    && for ver in ${helmversions[*]}; do wget --directory-prefix=/var/tmp/bin/${ver} https://storage.googleapis.com/kubernetes-helm/helm-${ver}-linux-amd64.tar.gz; done

# application image
FROM centos:7.4.1708

RUN yum -y install epel-release && \
    yum -y update && \
    yum -y install createrepo yum-utils nginx && \
    yum clean all

RUN mkdir /yum && \
    chmod 777 /yum && \
    mkdir -p /static && \
    chmod 777 /static

COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=downloader /var/tmp/rpm/ /yum
COPY --from=downloader /var/tmp/yaml/ /static
COPY --from=downloader /var/tmp/bin/ /static

COPY entrypoint.sh /entrypoint.sh
RUN chmod 700 /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]