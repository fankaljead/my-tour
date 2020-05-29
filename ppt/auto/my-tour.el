(TeX-add-style-hook
 "my-tour"
 (lambda ()
   (TeX-add-to-alist 'LaTeX-provided-class-options
                     '(("beamer" "14pt" "notheorems")))
   (TeX-add-to-alist 'LaTeX-provided-package-options
                     '(("ctex" "UTF8" "noindent")))
   (TeX-run-style-hooks
    "latex2e"
    "beamer"
    "beamer10"
    "ctex"
    "listings")
   (TeX-add-symbols
    '("mywarn" 1)
    '("mybold" 1)
    '("mylead" 1)
    "xiaoer"
    "xinwei"
    "emptypar")
   (LaTeX-add-environments
    "theorem"
    "definition"
    "example"))
 :latex)

