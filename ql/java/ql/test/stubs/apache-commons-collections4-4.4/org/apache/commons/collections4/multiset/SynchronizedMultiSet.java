// Generated automatically from org.apache.commons.collections4.multiset.SynchronizedMultiSet for testing purposes

package org.apache.commons.collections4.multiset;

import java.util.Set;
import org.apache.commons.collections4.MultiSet;
import org.apache.commons.collections4.collection.SynchronizedCollection;

public class SynchronizedMultiSet<E> extends SynchronizedCollection<E> implements MultiSet<E>
{
    protected SynchronizedMultiSet() {}
    protected MultiSet<E> decorated(){ return null; }
    protected SynchronizedMultiSet(MultiSet<E> p0){}
    protected SynchronizedMultiSet(MultiSet<E> p0, Object p1){}
    public Set<E> uniqueSet(){ return null; }
    public Set<MultiSet.Entry<E>> entrySet(){ return null; }
    public boolean equals(Object p0){ return false; }
    public int add(E p0, int p1){ return 0; }
    public int getCount(Object p0){ return 0; }
    public int hashCode(){ return 0; }
    public int remove(Object p0, int p1){ return 0; }
    public int setCount(E p0, int p1){ return 0; }
    public static <E> SynchronizedMultiSet<E> synchronizedMultiSet(MultiSet<E> p0){ return null; }
}
