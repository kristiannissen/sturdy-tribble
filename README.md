# Sturdy Tribble
## Sjov med ord i Go Lang

Som Corona-hobby-projekt syntes jeg det kunne være sjovt at implementere en dansk stemmer og fandt i den forbindelse denen side http://snowball.tartarus.org/algorithms/danish/stemmer.html.

Jeg er dog ikke 100% enig i betragtningerne beskrevet i dokumentet, f.eks endelsen "else" er ganske almindelig forkommende i det danske sprog. Så hvorfor ikke tage den med i main_suffix?
I stedet fjernes "e" og senere "els".

Efterfølgende lavede jeg Stopword baseret på en liste med danske stop ord jeg fandt her på github.

TODO: Brug funktions navne: main_suffix, consonant_pair, other_suffix, undouble

### Tf-idf

Det næste jeg vil forsøge mig med, er en term frequency-inverse document frequency løsning med udgangspunkt i http://www.tfidf.com/
