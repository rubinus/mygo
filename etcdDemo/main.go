package main

import (
	"time"

	"fmt"

	"context"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer cli.Close()

	//sr, err := cli.Status(context.TODO(), "localhost:2379")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("%+v\n", sr)

	le, err := cli.Lease.Grant(context.TODO(), 3)
	//fmt.Println(le.TTL)
	_, err = cli.KV.Put(context.TODO(), "/abc/def", "傻逼12", clientv3.WithLease(le.ID))
	//fmt.Println(pr, "test etcdDemo ...")
	//gr, err := cli.KV.Get(context.TODO(), "test")
	//fmt.Println(gr.Kvs)

	ctx, cancelFunc := context.WithCancel(context.TODO())
	defer cancelFunc()
	defer cli.Lease.Revoke(context.TODO(), le.ID)

	lech, err := cli.Lease.KeepAlive(ctx, le.ID)
	if err != nil {
		fmt.Println(err, "--keepalive")
	}

	kvc := clientv3.NewKV(cli)

	gr, err := kvc.Get(context.TODO(), "/abc/def", clientv3.WithPrevKV())
	fmt.Println("第一次get:", gr.Kvs, err, gr.Count)

	watcher := clientv3.NewWatcher(cli)
	wch := watcher.Watch(context.TODO(), "/abc/def", clientv3.WithPrevKV())
	go func() {
		for {
			select {
			case r := <-wch:
				fmt.Println("----watch---")
				fmt.Printf("%+v %s", r, "\n")
			case l := <-lech:
				if l == nil {
					fmt.Println("invalid keepalive", le.ID)
				} else {
					fmt.Println("-keepalive-", le.ID, l.ID, l.Revision)
				}
			}
		}
	}()
	//for x := range wch {
	//	for _, v := range x.Events {
	//		fmt.Printf("%+v\n", v)
	//	}
	//}

	txn := kvc.Txn(context.TODO())
	txn.If(clientv3.Compare(clientv3.CreateRevision("/abc/def2"), "=", 0)).
		Then(clientv3.OpPut("/abc/def2", "10000", clientv3.WithLease(le.ID))).
		Else(clientv3.OpGet("/abc/def2"))
	txnres, err := txn.Commit()
	if !txnres.Succeeded {
		fmt.Printf("锁占用")
	} else {
		fmt.Println("txnres", txnres)
	}

	time.Sleep(3500 * time.Millisecond)

	//fmt.Println("/n======")
	//gr, err = kvc.Get(context.TODO(), "/abc/def", clientv3.WithPrevKV())
	//fmt.Println(gr.Kvs, err, gr.Count)

}
