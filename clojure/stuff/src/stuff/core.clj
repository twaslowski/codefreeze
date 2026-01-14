(ns stuff.core
  (:gen-class))

(defn greet [name]
  (str "Hello, " name "!")
)

(defn is-even [x]
  (= (mod x 2) 0)
)

(defn filter-even [lst]
  (filter is-even lst)
)

(defn square [x]
    (* x x)
)

(defn -main
  [& args]
  (println
    (map square
      (filter is-even
        {2 "a" 4 "c"}
      )
    )
  )
)
