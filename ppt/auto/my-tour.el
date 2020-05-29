(TeX-add-style-hook
 "my-tour"
 (lambda ()
   (TeX-add-to-alist 'LaTeX-provided-class-options
                     '(("beamer" "14pt" "notheorems")))
   (TeX-add-to-alist 'LaTeX-provided-package-options
                     '(("ctex" "UTF8" "noindent")))
   (add-to-list 'LaTeX-verbatim-environments-local "lstlisting")
   (add-to-list 'LaTeX-verbatim-environments-local "semiverbatim")
   (add-to-list 'LaTeX-verbatim-macros-with-braces-local "lstinline")
   (add-to-list 'LaTeX-verbatim-macros-with-braces-local "href")
   (add-to-list 'LaTeX-verbatim-macros-with-braces-local "hyperref")
   (add-to-list 'LaTeX-verbatim-macros-with-braces-local "hyperimage")
   (add-to-list 'LaTeX-verbatim-macros-with-braces-local "hyperbaseurl")
   (add-to-list 'LaTeX-verbatim-macros-with-braces-local "nolinkurl")
   (add-to-list 'LaTeX-verbatim-macros-with-braces-local "url")
   (add-to-list 'LaTeX-verbatim-macros-with-braces-local "path")
   (add-to-list 'LaTeX-verbatim-macros-with-delims-local "lstinline")
   (add-to-list 'LaTeX-verbatim-macros-with-delims-local "path")
   (TeX-run-style-hooks
    "latex2e"
    "beamer"
    "beamer10"
    "ctex"
    "listings"
    "graphicx"
    "float"
    "subfigure")
   (TeX-add-symbols
    '("mywarn" 1)
    '("mybold" 1)
    '("mylead" 1)
    "xiaoer"
    "xinwei"
    "emptypar")
   (LaTeX-add-labels
    "Fig.system")
   (LaTeX-add-amsthm-newtheorems
    "theorem"
    "definition"
    "example"))
 :latex)

