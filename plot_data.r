chan <- read.csv("./output/chan.csv")
lock <- read.csv("./output/lock.csv")
dat <- list(a=unlist(chan["ns_op"][,1]),b=unlist(lock["ns_op"][,1]))
plot(unlist(dat),type="l",xlim=c(1,max(sapply(dat,length))), xlab = "Read percentage", ylab = "ns/op")
mapply(lines,dat,col=seq_along(dat),lty=1)
legend("topleft",c("Channel", "Lock"),lty=1,col=seq_along(dat))
