(ns sort)

(defn split-min
	[coll]
	(loop [m (first coll) o (rest coll) t []]
		(if (seq o)
			(let [n (first o)]
				(if (< n m)
					(recur n (rest o) (conj t m))
					(recur m (rest o) (conj t n))))
			[m t])))

(defn select-sort
	[coll]
	(if (<= (count coll) 1)
		coll
	(lazy-seq
		(let [[m o] (split-min coll)]
			(cons m (select-sort o))))))

(defn merge-coll
	[c1 c2]
	(loop [a c1 b c2 n []]
		(if (seq a)
			(if (seq b)
				(if (> (first a) (first b))
					(recur a (rest b) (conj n (first b)))
					(recur (rest a) b (conj n (first a))))
				(concat n a))
			(concat n b))))

(defn half
	[coll]
	(let [n (int (/ (count coll) 2))]
		(split-at n coll )))

(defn merge-sort
	[coll]
	(if (<= (count coll) 1)
		coll
		(let [[m n] (half coll)]
		  (merge-coll (merge-sort m) (merge-sort n)))))