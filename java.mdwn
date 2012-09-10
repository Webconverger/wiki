We discourage the use of Java on the Web for several reasons including
security, however some **legacy** Web applications unfortunately require Java.
Such as [Norway's banking applications](http://no.wikipedia.org/wiki/BankID).

In the [[chroot]], simply `apt-get update && apt-get install sun-java6-plugin`,
should prompt you to accept their license terms.

[[Test|testing]] Java support by visiting: <http://java.com/en/download/testjava.jsp>
