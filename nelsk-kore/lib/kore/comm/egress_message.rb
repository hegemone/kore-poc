module Kore
  module Comm
    class EgressMessage
      attr_reader :originator, :content
      def initialize(o)
        @originator = o.fetch(:originator, '')
        @content= o.fetch(:content, '')
      end

      def to_s
        "[#{@originator}] [#{@content}]"
      end
    end
  end
end
