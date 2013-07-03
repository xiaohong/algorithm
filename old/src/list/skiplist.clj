(ns core
  (:require [clojure.string :as string])
  (use [clojure.stacktrace]))

(def ^:dynamic *max-level* 7)

(defrecord Node [x forward])

(defn new-node
  [x level]
  (Node. x  (make-array (class (Object.)) (inc level))))

(defrecord SkipList [header level])

(defn new-skiplist
  []
  (SkipList. (new-node nil *max-level*) 0))

(defn insert
  [sl v]
  (let [size (.level sl)
        h (.header sl)
        [update x] (loop [s size update [] h h]
                     (if (>= s 0)
                       (let [n (loop [x h]
                                 (if (and (not (nil? (aget (.forward x) s)))
                                          (< (.x (aget (.forward x) s)) v))
                                   (recur (aget (.forward x) s))
                                   x))]
                         (recur (dec s) (cons n update) n))
                         [update h]))
        level (rand-level)
        nn (new-node v level)
        [size update] (if (<= level size)
                         [size update]
                         [level (loop [i (inc size) update (vec update)]
                                  (if (<= i level)
                                    (recur (inc i) (conj  update h))
                                    update))])]
    (loop [ l 0]
      (if (<= l level)
        (do
          (aset (.forward nn) l (aget (.forward (nth update l)) l))
          (aset (.forward (nth update l)) l nn)
          (recur (inc l)))
        (SkipList. h size)))))

(defn- rand-level
  []
  (let [l (/ (Math/log (- 1.0 (Math/random)))
             (Math/log (- 1.0 0.5)))]
    (min *max-level* (int l))))


(defn not-nil?
  [n]
  (not (nil? n)))

(defn delete 
  [sl v])

(defn print-skiplist
  [sl]
  (loop [h (.header sl) n []]
    (if (not-nil? (aget (.forward h) 0))
      (recur (aget (.forward h) 0)
             (conj n (.x (aget (.forward h) 0))))
      n)))

(defn test-list
  []
  (let [t (new-skiplist)
        t (loop [i 0 t t]
            (if (< i 100) 
              (do
                (recur (inc i) (insert t (rand-int 100))))
              t))]
    (println (print-skiplist t))
    ))

(test-list)
